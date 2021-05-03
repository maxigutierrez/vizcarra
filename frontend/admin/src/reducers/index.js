import { combineReducers } from 'redux';
import { routerReducer } from 'react-router-redux';
import settings from './settings';
import user from './user';

const rootReducer = combineReducers({
  settings: settings,
  user: user,
  routing: routerReducer
});

export default rootReducer;
