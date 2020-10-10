import axios from "axios";
import {
  GET_POSTS,
  GET_POST,
  ADD_POST,
  EDIT_POST,
  UPDATE_LIKES,
  POST_ERROR,
  ADD_COMMENT,
  REMOVE_COMMENT,
  REMOVE_POST
} from "./types";

import setAuthToken from "../../utils/setAuthToken";

//get all posts
export const getPosts = () => async dispatch => {
  try {
    const res = await axios.get("/api/posts");
    dispatch({
      type: GET_POSTS,
      payload: res.data
    });
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: POST_ERROR
    });
  }
};

//add a post
export const addPost = (formData, history) => async dispatch => {
  try {
    setAuthToken(localStorage.token);
    const config = {
      headers: {
        "Content-Type": "application/json"
      }
    };
    const res = await axios.post("/api/posts", formData, config);
    dispatch({
      type: ADD_POST,
      payload: res.data
    });
    history.push("/posts");
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: POST_ERROR
    });
  }
};

//remove a post
export const removePost = post_id => async dispatch => {
  try {
    const res = await axios.delete(`/api/posts/${post_id}`);
    dispatch({
      type: REMOVE_POST,
      payload: post_id
    });
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: POST_ERROR
    });
  }
};

//add a like
export const like = id => async dispatch => {
  try {
    const res = await axios.put(`/api/posts/likes/${id}`);
    dispatch({
      type: UPDATE_LIKES,
      payload: { id, likes: res.data }
    });
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: POST_ERROR
    });
  }
};

//add comment
export const comment = (text, id) => async dispatch => {
  try {
    setAuthToken(localStorage.token);
    const config = {
      headers: {
        "Content-Type": "application/json"
      }
    };
    const res = await axios.post(`/api/posts/comments/${id}`, text, config);
    dispatch({
      type: ADD_COMMENT,
      payload: { id: id, comments: res.data }
    });
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: POST_ERROR
    });
  }
};

//remove comment
export const removeComment = (post_id, comment_id) => async dispatch => {
  try {
    setAuthToken(localStorage.token);
    const res = await axios.delete(
      `/api/posts/comments/${post_id}/${comment_id}`
    );
    dispatch({
      type: REMOVE_COMMENT,
      payload: { postId: post_id, comments: res.data }
    });
    console.log("from @comments", res);
  } catch (error) {
    console.error(error.message);
    dispatch({
      type: POST_ERROR
    });
  }
};
