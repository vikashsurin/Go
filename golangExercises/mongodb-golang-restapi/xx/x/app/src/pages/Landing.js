import React from "react";
import { connect } from "react-redux";

const Landing = props => {
  return (
    <div className='landing'>
      <div>
        <span>T</span>
        <span>
          hroughout years the way of doing things have advanced themselves
          drastically , whether it be medical science , technical science and
          the list goes onn....yes, I mean the internet is doing great wonders .
        </span>
      </div>
      <div>
        <span>D</span>
        <span>
          o you think the way of learning should have a positive effect of the
          internet , here we begin - just let’s give it a try ...
        </span>
      </div>
      <div>
        <span>T</span>
        <span>
          he concept of the app is based on a very simple yet effective action :
          “Debate” - check it out .
        </span>
      </div>
      <div className='footer'>
        <a href=''>link</a>
      </div>
    </div>
  );
};

const mapStateToProps = state => {
  return state;
};

export default connect(mapStateToProps)(Landing);
