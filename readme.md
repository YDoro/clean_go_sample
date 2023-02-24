# Project TLDR
- DomainProduct exports ProductRepository interface

- PostgresProductRepository implement it 

- DomainProductUseCase expects a ProductRepository instance

- AdapterHttpProductService expects a usecase Instance

- DIProduct create a new PostgresProductRepository

- then pass it to a instance of the Domain/ProductUseCase

- then pass it to a instance of the ApdapterHttpProductService

- the main uses DIConfigProduct as the product service and pass it to the http handler

- DTO's are used to make the entire transaction


# What I've learned with this project

##  Folder structure  

My folder structure in other projects are very messy and it was causing cyclic dependency error.

The DI folder is a great thing that i haven't thinked, it makes easier to read the conde on `main.go`

The DTO folder outside the domain layer makes more sense than have all together as I have in another projects

### The Adapter folder containing things of the infrastructure layer makes me think about it...
    -   adapter should adapt an external dependency to avoid heavy coupling
    -   infra should connect adapters with the buisiness rules
    -   here the adapters are making the two roles


## Files
    - my files in other projects have too many lines instead several short files

## Packages 
 my packages in other projects have naming issues, here we can see names like `productrepository` or `productusecase` which makes the code easier to read and more "clean architecture" way

## Interface segregation
In other projects the interface segregation are most improved than at this one, by example the domain/ProductRepository has two methods used in diffrent usecases, here we can split it into AddProductRepository and FetchProduct repository to respect the ISP and make the tests easier
