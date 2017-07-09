# gSocialAPI
**gSocialAPI** is simple microservice which provide social endpoints for implementing social activities in your application.

Available functions:

*	Following
* 	Favorites


### Implementation 
**gSocialAPI** is born with refactoring social activities api implemented as part of one project in PHP / Codeigniter, pretty monolithic.
Following microservice design patterns most of the social interaction that was implemented in PHP using Codeigniter MVC framework was reimplemented as Go microservice using in the same time master database MySQL RDBMS and Redis the most popular NoSQL. This microservice is basic implementation of the real service used in the project as example for how to implement microservice for extending social interaction in the application using in the same time RDBMS for data storage and NoSQL for social activities tracking. This microservice will be one of the following used as building blocks for building modern backend architecture.

### Why MySQL and Redis and not MongoDB
The answer is simple, we love SQL :) MySQL RDBMS is already tested, proved and implemented as database engine in the backend in other projects also is much easier to find great RDBMS admin for cheap. Redis on the other side is super fast NoSQL database based on data structures which simplify the logic for tracking social connection,  tracking favorite products or other social activities.
MongoDB is well used in a lot of cases as DB engine in the backend mainly because of the easy horizontal scaling and developers friendly api but on the other side you must change your mindset for designing the application architecture based on documents and not tables which also require time and experience. Implementing MongoDB would require changing most of the logic at the backend modules and testing how will fit at the application concept. So the choice at the end was very simple, using MySQL RDBMS for storing the data which is not often modified and is used most of the time for reading and Redis NoSQL for storing dynamic data which require fast r/w and access to the specific information.

Database structure will not be presented but following the simple SQL query its easy to adapt for other cases.

