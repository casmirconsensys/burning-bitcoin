## TODOs:


## How to add a view function
1. Add necessary getter(s) in `x/burning/keeper/keeper.go`
1. Add response type to `x/burning/types/querier.go`
    1. Response type is a struct with the return values
    1. Implement `String()` for the response type
1. Add function to querier `x/burning/keeper/querier.go`
    1. Add new string tag for the new query
    1. Add new `query___` function
    1. Add new case block to `switch` in `NewQuerier()`
1. Add alias type in `x/burning/alias.go`
1. Add to CLI  
    1. add to `x/burning/client/cli/query.go`
1. Add to REST
    1. add to `x/burning/client/rest/query.go`
    1. add GET route to `x/burning/client/rest/rest.go`

## How to add a non-view function
1. Add necessary getters/setters in `x/burning/keeper/keeper.go`
1. Add msg type in `x/burning/types/msgs.go`
    1. Message type is a struct with the arguments
    1. Implement `New___()`
    1. Implement `GetSigners()`
    1. Implement `Type()`
    1. Implement `ValidateBasic()`
    1. Implement `GetSignBytes()`
    1. Implement `Route()`
1. Add to handler
    1. Add new `handle____` function
    1. Add new case block to `switch` in `NewHandler()`
1. Add aliases in `x/burning/alias.go`
    1. Add alisa in `var` block
    1. Add alias in `type` block
1. Add to CLI  
    1. add to `x/burning/client/cli/tx.go`
1. Add to REST
    1. add to `x/burning/client/rest/tx.go`
    1. add POST route to `x/burning/client/rest/rest.go`
