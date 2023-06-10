import { ApolloClient, InMemoryCache, createHttpLink } from "@apollo/client";
import { setContext } from '@apollo/client/link/context';
import { BASE_URL, TOKEN } from './utils/constants';

const httpLink = createHttpLink({
  uri: `${BASE_URL}/graphiql`,
});

const authLink = setContext((_, { headers }) => {
  const token = localStorage.getItem(TOKEN);

  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    }
  }
});

const client = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
});

export default client;