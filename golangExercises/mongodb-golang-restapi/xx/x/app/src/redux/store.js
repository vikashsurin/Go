import { createStore, applyMiddleware } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';
import thunk from 'redux-thunk';
import rootReduer from './reducers';

const initialState = {};
const middlewares = [thunk];

const store = createStore(
  rootReduer,
  initialState,
  composeWithDevTools(applyMiddleware(...middlewares))
);

export default store;
