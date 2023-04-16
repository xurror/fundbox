import { ApolloClient, InMemoryCache } from "@apollo/client";

const client = new ApolloClient({
    uri: "https://countries.trevorblades.com", // Dummy data gotten from country api
    cache: new InMemoryCache(),
});

export default client;