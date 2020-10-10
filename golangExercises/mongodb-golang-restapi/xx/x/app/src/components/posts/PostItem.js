import React, { Fragment, useState, useEffect } from "react";
import { connect } from "react-redux";
import { withRouter } from "react-router-dom";
import { like } from "../../redux/actions/@posts";
import PropTypes from "prop-types";
import Comment from "./Comment";
import CommentForm from "./CommentForm";
import { UPDATE_LIKES } from "../../redux/actions/types";

const PostItem = ({
  auth,
  history,
  user,
  like,
  post: { _id, title, text, name, avatar, date, comments, likes }
}) => {
  const togLike = user ? likes.filter(like => like.user === user._id) : "null";

  const likePost = e => {
    auth ? like(_id) : history.push("/login");
  };
  const [commentToggler, setCommentToggler] = useState(false);
  const [commentFormToggler, setCommentFormToggler] = useState(false);
  const toggleComments = () => {
    setCommentToggler(!commentToggler);
  };

  return (
    <Fragment>
      <div className='pc-wrapper'>
        <div className='pc-top'>
          <img height={"50px"} src={avatar} alt='pic' />
          <span className='title'>{title}</span>
          <div className='auth-dat'>
            <span className='author'>
              <i>{name}</i>
            </span>
            <span className='date'>{new Date(date).toLocaleString()}</span>
          </div>
        </div>
        <div className='pc-body'>{text}</div>

        <div className='pc-counts'>
          <span className='li'>
            likes:{likes.length > 0 ? likes.length : ""}
          </span>
          <span
            className='co'
            onClick={() => [
              setCommentToggler(!commentToggler),
              setCommentFormToggler(!commentFormToggler)
            ]}
          >
            comments:{comments.length > 0 ? comments.length : ""}
          </span>
        </div>

        {auth && (
          <div className='pc-buttons'>
            <div className='buttons' onClick={e => likePost(e)}>
              <span>
                {togLike[0] === undefined || togLike.toString() === "null"
                  ? "like"
                  : auth === true
                  ? "liked"
                  : "like"}
              </span>
            </div>
            <div
              className='buttons'
              onClick={() => setCommentFormToggler(!commentFormToggler)}
            >
              comment
            </div>
            <div className='buttons'>share</div>
          </div>
        )}
        {commentToggler && auth && (
          <Comment auth={auth} admin={user} post_id={_id} comments={comments} />
        )}
        {commentFormToggler && auth && <CommentForm id={_id} />}
      </div>
    </Fragment>
  );
};
PostItem.propTypes = {
  like: PropTypes.func.isRequired,
  toggleComments: PropTypes.func.isRequired
};
const mapStateToProps = state => {
  return { posts: state.posts };
};
export default connect(mapStateToProps, { like })(withRouter(PostItem));
