## Monolith are hard to scale
- Monoliths are easy to develop but hard to scale, when the codebase becomes large.
- when multiple people need to work on it, there is a lot of dependency

## Loosely coupled
- Microservices, is an assembly between loosely coupled services.
- each service doesn't need to know too much about another service.
- the services are independently deployable, scalable, maintainable and testable.

## Seperate db's
- each service maintains a seperate db, even if the seperation is only logical
- the schema is private to only one service
- you don't share schema between two services

## Domain driven design
- User interface - application layer - domain layer - infrastructure layer

## Sagas
- Using the SAGA pattern, it helps you maintain data consistency across services.
- More like a transaction, where you chain sequence of steps accross microservices until you eventually complete a microservice

## Books
- building microservices
- domain-driven microservices
- microservices pattern