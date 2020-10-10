import React, { useState } from "react";
import axios from "axios";

function Form() {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });
  const { email, password } = formData;
  const onChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  /// fetch
  const onSubmit = async (e) => {
    e.preventDefault();
    const url = "http://localhost:8080/login";
    // const options = {
    //   method: "POST",
    //   headers: {
    //     Accept: "application/json",
    //     "Content-Type": "application/json;charset=UTF-8",
    //   },
    //   body: JSON.stringify({
    //     email: email,
    //     password: password,
    //   }),
    // };

    // fetch(url, options).then((response) => {
    //   console.log(response);
    // });

    const res =await  axios({
      method: "post",
      url: "http://localhost:8080/login",
      data: {
        email: "vfdf",
        password: "sdfdsf",
      },
    });
    console.log(res.data.Email,res.data.Password);
  };
  return (
    <div className="Form">
      <form onSubmit={(e) => onSubmit(e)}>
        <input
          type="text"
          name="email"
          value={email}
          onChange={(e) => onChange(e)}
          placeholder="email"
        />

        <input
          type="text"
          name="password"
          value={password}
          onChange={(e) => onChange(e)}
          placeholder="password"
        />
        <input type="submit" />
      </form>
    </div>
  );
}

export default Form;
