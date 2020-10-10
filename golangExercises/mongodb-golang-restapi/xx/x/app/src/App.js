import React, { Fragment, useEffect } from "react";
import { Provider } from "react-redux";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
// import "./styles/main.scss";

//componenets
import PrivateRoute from "./components/PrivateRoute";
import Navbar from "./components/Navbar";
import Landing from "./pages/Landing";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Dashboard from "./pages/Dashboard";
import Posts from "./pages/Posts";
import CreateProfile from "./components/profile/CreateProfile";
import AddPost from "./components/posts/AddPost";
import Auth from "./pages/Auth";
//redux
import store from "./redux/store";
import { loadUser } from "./redux/actions/@auth";
import EditProfile from "./components/profile/EditProfile";

const App = props => {
  useEffect(() => {
    store.dispatch(loadUser());
  }, []);
  return (
    <Fragment>
      <Provider store={store}>
        <Router>
          <Navbar />
          <Switch>
            <Route exact path='/' component={Landing} />
            <Route exact path='/login' component={Login} />
            <Route exact path='/register' component={Register} />
            <Route exact path='/auth' component={Auth} />
            <Route exact path='/logout' component={Auth} />
            <Route exact path='/posts' component={Posts} />
            <Route exact path='/create-profile' component={CreateProfile} />
            <Route exact path='/edit-profile' component={EditProfile} />
            <PrivateRoute exact path='/dashboard' component={Dashboard} />
            {/* <PrivateRoute exact path='/add-post' component={AddPost} /> */}
          </Switch>
        </Router>
      </Provider>
    </Fragment>
  );
};

// App.propTypes = {};

export default App;
