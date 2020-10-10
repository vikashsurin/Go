import React, { useState, Fragment } from "react";
import { connect } from "react-redux";
import { withRouter } from "react-router-dom";
import PropTypes from "prop-types";
import { addPost } from "../../redux/actions/@posts";

const AddPost = ({ addPost, history }) => {
  const [formData, setFormData] = useState({
    title: "",
    text: ""
  });
  const [showForm, setShowForm] = useState(false);
  const { title, text } = formData;
  const onChange = e => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };
  const onSubmit = e => {
    e.preventDefault();
    addPost(formData, history);
    setShowForm(!showForm);
    setFormData({ title: "", text: "" });
  };

  return (
    <Fragment>
      <form className='addpost' onSubmit={e => onSubmit(e)}>
        <input
          type='text'
          name='title'
          placeholder='   Title'
          autoComplete='off'
          value={title}
          onChange={e => onChange(e)}
          onClick={() => setShowForm(!showForm)}
        />
        <br />
        {showForm && (
          <Fragment>
            <textarea
              type='text'
              name='text'
              rows='20'
              placeholder='    about topic in detail'
              value={text}
              onChange={e => onChange(e)}
            />
            <br />
            <button type='submit'> create post</button>
          </Fragment>
        )}
      </form>
    </Fragment>
  );
};

AddPost.propTypes = {
  addPost: PropTypes.func.isRequired
};
export default connect(null, { addPost })(withRouter(AddPost));
