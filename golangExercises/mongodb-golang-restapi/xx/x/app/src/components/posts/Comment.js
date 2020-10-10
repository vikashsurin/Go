import React, { Fragment, useState, useEffect } from "react";
import { connect } from "react-redux";
import { removeComment } from "../../redux/actions/@posts";
import PropTypes from "prop-types";

const Comment = ({ admin, auth, post_id, removeComment, comments }) => {
  const commentList = comments
    ? comments.map(comment => {
        console.log(comment.name);
        return (
          <div className='comment'>
            <img
              className='img'
              src={comment.avatar}
              alt='pic'
              onClick={() => console.log(comment.avatar)}
            />
            <div className='detail'>
              <a href='#' className='name'>
                {comment.name}
              </a>
              <span className='text'>{comment.text.toString()}</span>
              <span className='date'>{comment.date}</span>
            </div>

            {auth && admin._id
              ? admin._id === comment.user && (
                  <div
                    className='dot'
                    onClick={() => removeComment(post_id, comment._id)}
                  >
                    ...
                  </div>
                )
              : ""}
          </div>
        );
      })
    : "";
  return (
    <Fragment>
      <div className='c-container'>{commentList}</div>
    </Fragment>
  );
};

Comment.propTypes = {
  admin: PropTypes.object.isRequired,
  auth: PropTypes.bool.isRequired,
  post_id: PropTypes.object.isRequired,
  removeComment: PropTypes.func.isRequired,
  comments: PropTypes.object.isRequired
};
const mapStateToProps = state => {
  return { state };
};
export default connect(mapStateToProps, { removeComment })(Comment);
