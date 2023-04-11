import React from 'react';
import { useNavigate } from 'react-router-dom'
import "../stylessheets/Login.css";
import 'bootstrap/dist/css/bootstrap.min.css';

export default function Login() {
  const [username, setUsername] = React.useState('')
  const [password, setPassword] = React.useState('')

  const navigate = useNavigate()
  // fucnion para capturar los valores de los input usuario y password
  // puede tener cualquier nombre

  const chageValues = () => {
    setUsername('')
    setPassword('')
  }

  const onSubmit = (e) => {
    e.preventDefault()
    if (username === "erick" && password === "123") {
      console.log('Entre')
      navigate('/seleccion')
    } else {
      navigate('/')
      alert("Credenciales erroreas")
    }
  }

  return (
    < div >
      {chageValues}
      <div className="login">
        <form  >
          <img src={require("../imagenes/login2.png")} className="rounded mx-auto d-block" alt="..." />
          <div className="mb-3">
            <label className="form-label">Usuario</label>
            <input
              type="text"
              className="form-control"
              value={username}
              placeholder='Username'
              onChange={e => setUsername(e.target.value)}
            />
          </div>
          <div className="mb-3">
            <label className="form-label">Contraseña</label>
            <input
              type="password"
              placeholder='Password'
              className="form-control"
              value={password}
              onChange={e => setPassword(e.target.value)}
            />
          </div>
          <div className="mb-3">
            <button type="submit" className="btn btn-primary" onClick={onSubmit}>
              Login
            </button>
          </div>
          
        </form>
      </div >
    </div >
  )
}