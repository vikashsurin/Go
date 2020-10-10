import { combineReducers } from "redux";
import auth from "./_auth";
import profile from "./_profile";
import posts from "./_posts";

export default combineReducers({
  auth,
  profile,
  posts
});
