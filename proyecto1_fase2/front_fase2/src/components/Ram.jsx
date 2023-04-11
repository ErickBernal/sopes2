import { useNavigate } from 'react-router-dom'
import React, { useState, useEffect } from "react"
import axios from 'axios'
import { Pie } from 'react-chartjs-2'
import PiesChart from './PiesChart'

export default function Ram() {
    const navigate = useNavigate()
    const [total, setTotal] = useState(0);
    const [Disponible, setDisponible] = useState(0);
    const [En_uso, setEn_uso] = useState(0);
    const [usb, setUsb] = useState(false);
    const [textoBoton, setTextoBoton] = useState("");
    const [isToggled, setIsToggled] = useState(false);



    const getRam = () => {
        axios.get('http://localhost:8080/ram')
            .then(response => {
                // Aquí puedes manejar la respuesta
                console.log(response.data);
                setTotal(response.data.total);
                setDisponible(response.data.disponible);
                setEn_uso(response.data.usado);
            })
            .catch(error => {
                // Aquí puedes manejar el error
                console.error(error);
            });
    }

    useEffect(() => {
        const intervalId = setInterval(() => {
            getRam();
        }, 1000);
        return () => clearInterval(intervalId);
    }, []);

    const toggle = () => {
        setUsb(!usb);
    }

    const handleClick = () => {
        if (isToggled) {
            // Ejecuta la función 1
            console.log('Función 1');
        } else {
            // Ejecuta la función 2
            console.log('Función 2');
        }

        setIsToggled(!isToggled);
    }


    return (
        <div className="login">
            <form>
                <div className='row'>
                    <div className='col'>
                        <p className="form-label"><b>Memoria Ram: </b>Grafico en tiempo real</p>
                        <div className="bg-light mx-auto border border-2 border-primary" style={{ width: "450px", height: "250px" }}>
                            <div style={{ width: "100%", height: "100%", padding: "10px 0" }}>
                                <PiesChart
                                    total={total}
                                    disponible={Disponible}
                                    en_uso={En_uso}
                                />
                            </div>
                            <p className="form-label"><b> Total: </b>{total.toFixed(2)} GB</p>
                            <p className="form-label"><b> Disponible: </b>{Disponible.toFixed(2)} GB</p>
                            <p className="form-label"><b> En_uso: </b>{En_uso.toFixed(2)} GB</p>
                            <button
                                className="btn btn-primary"
                                onClick={() => { navigate("/") }}
                            >
                                Login
                            </button>
                            <button type="submit" className="btn btn-secondary" onClick={() => { navigate("/seleccion") }}>
                                Seleccion
                            </button>
                        </div>

                    </div>

                </div>
            </form>
        </div>
    );
}
