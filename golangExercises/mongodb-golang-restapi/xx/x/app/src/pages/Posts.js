import React, { Fragment, useEffect } from "react";
import { connect } from "react-redux";

import { Link } from "react-router-dom";
import PropTypes from "prop-types";
import { getPosts } from "../redux/actions/@posts";
import PostItem from "../components/posts/PostItem";
import AddPost from "../components/posts/AddPost";
import PostPlaceholder from "../components/placeholder/PostPlaceholder";

const Posts = ({
  getPosts,
  posts: { posts, isLoading },
  auth: { isAuthenticated, user }
}) => {
  useEffect(() => {
    getPosts();
  }, []);

  const postsList = isLoading ? (
    <Fragment>
      <PostPlaceholder />
      <PostPlaceholder />
      <PostPlaceholder />
    </Fragment>
  ) : (
    posts.map(post => {
      return (
        <Fragment key={post._id}>
          <PostItem
            key={post._id}
            auth={isAuthenticated}
            user={user}
            post={post}
          />
        </Fragment>
      );
    })
  );
  return (
    <Fragment>
      {/* <Link to='/add-post'>add Post</Link> */}
      <div class='main-body'>
        {isAuthenticated && <AddPost />}
        <div className='posts'>{postsList}</div>
      </div>
    </Fragment>
  );
};

Posts.propTypes = {
  getPosts: PropTypes.func.isRequired,
  posts: PropTypes.object.isRequired,
  isAuthenticated: PropTypes.bool.isRequired,
  toggleComState: PropTypes.object.isRequired
};

const mapStateToProps = state => {
  return {
    posts: state.posts,
    auth: state.auth
  };
};
export default connect(mapStateToProps, { getPosts })(Posts);
