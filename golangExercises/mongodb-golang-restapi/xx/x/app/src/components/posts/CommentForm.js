import React, { Fragment, useState } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { comment } from "../../redux/actions/@posts";

const CommentForm = ({ id, comment }) => {
  const [text, setText] = useState("");
  const onSubmit = e => {
    e.preventDefault();
    comment({ text }, id);
    setText("");
  };

  return (
    <Fragment>
      <form className='cf-form' onSubmit={e => onSubmit(e)}>
        <div className='cf-div'>
          <input
            className='cf-input'
            type='text'
            name='text'
            value={text}
            onChange={e => setText(e.target.value)}
            placeholder='write a comment'
          />
          <button type='submit'>send</button>
        </div>
      </form>
    </Fragment>
  );
};

CommentForm.propTypes = {
  comment: PropTypes.func.isRequired
};

export default connect(null, { comment })(CommentForm);
