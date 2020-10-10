import React,{useState} from 'react';
import axios from "axios"

function Form() {
  const [formData,setFormData] = useState({
    email:"",
    password:""
  });
const {email,password} = formData;
const onChange = e => {
  setFormData({ ...formData, [e.target.name]: e.target.value });
};
const onSubmit = async e => {
  e.preventDefault();
  setFormData({ title: "", text: "" });
  // const config = {
  //   headers: {
  //     "Content-Type": "application/json"
  //   }
  // }; 
  //  const res = await axios({
  //   method: 'post',
  //   url: 'http://localhost:8080/login',
  //   data: {
  //     "email": 'vikash@gmail.com',
  //     "password": 'mypassword'
  //   }
  // });
  const res = await axios.post('http://localhost:8080/login', {
    "email":"vikash@gmail.com",
    "password": "mypassword"
  })
  
  console.log(res.data)
};

  return (
    <form className="form" onSubmit={e=>onSubmit(e)}>
        <input type="text" name="email" value={email} onChange={e=>onChange(e)}/>
        <input type="text" name="password"  value={password} onChange={e=>onChange(e)}/>
        <input type="submit"/>
    </form>
  );
}

export default Form;
