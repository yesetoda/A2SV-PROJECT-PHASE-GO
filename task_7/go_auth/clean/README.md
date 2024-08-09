After reading the uncle Bob’s Clean Architecture Concept, I’m trying to implement it in Golang. This is a similar architecture that we used in our company, Kurio - App Berita Indonesia, but with a little different structure. Not too different, same concept but different in the folder structure.

You can look for a sample project here https://github.com/bxcodec/go-clean-arch, a sample CRUD management article.


Disclaimer :
I’m not recommending any library or framework used here. You could replace anything here, with your own or third party that has the same functions.
Basic
As we know the constraint before designing the Clean Architecture are :

Independent of Frameworks. The architecture does not depend on the existence of some library of feature-laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
Independent of any external agency. In fact, your business rules simply don’t know anything at all about the outside world.
More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

So, based on this constraint, every layer must independent and testable.

If Uncle Bob’s Architecture, has 4 layers:

Entities
Usecase
Controller
Framework & Driver
In my projects, I’m using 4 too :

Models
Repository
Usecase
Delivery
Models
Same as Entities, it will be used in all layers. This layer will store any Object’s Struct and its method. Example: Article, Student, Book.
Example struct :


Any entities or models will be stored here.

Repository
The repository will store any Database handler. Querying, or Creating/ Inserting into any database will be stored here. This layer will act for CRUD to the database only. No business process happens here. Only plain function to Database.

This layer also has the responsibility to choose what DB will use in the Application. Could be Mysql, MongoDB, MariaDB, Postgresql whatever, will decide here.

If using ORM, this layer will control the input, and give it directly to ORM services.

If calling microservices, it will be handled here. Create HTTP Requests for other services, and sanitize the data. This layer must fully act as a repository. Handle all data input-output no specific logic happen.

This Repository layer will depend on Connected DB, or other microservices if exists.

Usecase
This layer will act as the business process handler. Any process will be handled here. This layer will decide, which repository layer will use. And have the responsibility to provide data to serve into delivery. Process the data doing calculation or anything will be done here.

Usecase layer will accept any input from the Delivery layer, that is already sanitized, then process the input could be storing into DB, or Fetching from DB, etc.

This Usecase layer will depend to Repository Layer

Delivery
This layer will act as the presenter. Decide how the data will be presented. Could be as REST API, HTML File, or gRPC whatever the delivery type.
This layer also will accept the input from the user. Sanitize the input and sent it to the Usecase layer.

For my sample project, I’m using REST API as the delivery method.
The client will call the resource endpoint over the network, and the Delivery layer will get the input or request, and send it to Usecase Layer.

This layer will depend on Usecase Layer.

Communications Between Layer
Except for Models, each layer will communicate through an interface. For example, the Usecase layer needs the Repository layer, so how do they communicate? The repository will provide an interface to be their contract and communication.

Example of Repository’s Interface


Usecase layer will communicate to Repository using this contract, and the Repository layer MUST implement this interface so can be used by Usecase

Example of Usecase’s Interface


Same with Usecase, the Delivery layer will use this contract interface. And Usecase layer MUST implement this interface.

Testing Each Layer
As we know, clean means independent. Each layer is testable even though other layers don’t exist yet.

Models Layer
This layer only tested if any function/method was declared in any of Struct.
And can test easily and independently to other layers.
Repository
To test this layer, the better way is to do Integration testing. But you also can do mocking for each test. I’m using github.com/DATA-DOG/go-sqlmock as my helper to mock query process msyql.
Usecase
Because this layer depends on the Repository layer, means this layer needs a Repository layer for testing. So we must make a mockup of the Repository
that mocked with mockery, based on the contract interface defined before.
Delivery
Same with Usecase, because this layer depends on the Usecase layer, means we need the Usecase layer for testing. And Usecase layer also must be mocked with mockery, based on the contract interface defined before
For mocking, I use mockery for golang by vektra could be seen here https://github.com/vektra/mockery

Repository Test
To test this layer, as I said before, I’m using a sql-mock to mock my query process. You can use like what I used here github.com/DATA-DOG/go-sqlmock , or another that has a similar function

func TestGetByID(t *testing.T) {
 db, mock, err := sqlmock.New() 
 if err != nil { 
    t.Fatalf(“an error ‘%s’ was not expected when opening a stub  
        database connection”, err) 
  } 
 defer db.Close() 
 rows := sqlmock.NewRows([]string{
        “id”, “title”, “content”, “updated_at”, “created_at”}).   
        AddRow(1, “title 1”, “Content 1”, time.Now(), time.Now()) 
 query := “SELECT id,title,content,updated_at, created_at FROM 
          article WHERE ID = \\?” 
 mock.ExpectQuery(query).WillReturnRows(rows) 
 a := articleRepo.NewMysqlArticleRepository(db) 
 num := int64(1) 
 anArticle, err := a.GetByID(num) 
 assert.NoError(t, err) 
 assert.NotNil(t, anArticle)
}
Usecase Test
Sample test for Usecase layer, that depends on Repository layer.


Mockery will generate a mockup of the repository layer for me. So I don’t need to finish my Repository layer first. I can work on finishing my Usecase first even though my Repository layer is not implemented yet.

Delivery Test
The delivery test will depend on how you deliver the data. If using HTTP REST API, we can use httptest a built-in package for httptest in golang.

Because it depends on Usecase, so we need a mock of Usecase. Same with Repository, I’m also using Mockery to mock my use case, for delivery testing.

func TestGetByID(t *testing.T) {
 var mockArticle models.Article 
 err := faker.FakeData(&mockArticle) 
 assert.NoError(t, err) 
 mockUCase := new(mocks.ArticleUsecase) 
 num := int(mockArticle.ID) 
 mockUCase.On(“GetByID”, int64(num)).Return(&mockArticle, nil) 
 e := echo.New() 
 req, err := http.NewRequest(echo.GET, “/article/” +  
             strconv.Itoa(int(num)), strings.NewReader(“”)) 
 assert.NoError(t, err) 
 rec := httptest.NewRecorder() 
 c := e.NewContext(req, rec) 
 c.SetPath(“article/:id”) 
 c.SetParamNames(“id”) 
 c.SetParamValues(strconv.Itoa(num)) 
 handler:= articleHttp.ArticleHandler{
            AUsecase: mockUCase,
            Helper: httpHelper.HttpHelper{}
 } 
 handler.GetByID(c) 
 assert.Equal(t, http.StatusOK, rec.Code) 
 mockUCase.AssertCalled(t, “GetByID”, int64(num))
}
Final Output and The Merging
After finishing all layers and already passed on testing. You should merge into one system in main.go in the root project.
Here you will define, and create every need for the environment, and merge all layers into one.

Look for my main.go as an example:


You can see, that every layer merges into one with its dependencies.

Conclusion :
In short, if drawn in a diagram, can seen below

Every library used here you can change on your own. Because the main point of clean architecture is: no matter your library, your architecture is clean, testable also independent
This is how I organized my projects, you could argue, or agree, or maybe improve this for a better one, just leave a comment and share this
The Sample Projects
The sample project can be seen here https://github.com/bxcodec/go-clean-arch

Libary used for my project:

Glide : for package management
go-sqlmock from github.com/DATA-DOG/go-sqlmock
Testify : for testing
Echo Labstack (Golang Web Framework) for Delivery layer
Viper : for environment configurations
Further Reading about Clean Architecture :

Second part of this article: https://hackernoon.com/trying-clean-architecture-on-golang-2-44d615bf8fdf
https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html
http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/ . Another version of Clean Architecture in Golang
If you have a question , or need more explanation, or something, that I can not explain well here, you can ask me from my linkedin or email me. Thank you