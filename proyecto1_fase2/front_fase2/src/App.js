import './App.css';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom"
import Login from './components/Login';
import Ram from './components/Ram';
import Seleccion from './components/Seleccion';
import Usb from './components/Usb';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="ram" element={<Ram />} />
        <Route path="seleccion" element={<Seleccion />} />
        <Route path="usb" element={<Usb />} />

      </Routes>
    </Router>
  );
}

export default App;
