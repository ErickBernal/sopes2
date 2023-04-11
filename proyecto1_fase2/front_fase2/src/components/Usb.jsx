import { useNavigate } from 'react-router-dom'
import React, { useState, useEffect } from "react"
import axios from 'axios'



export default function Usb() {
    const navigate = useNavigate()


    const onSubmit_seleccion = () => {
        navigate('/seleccion')
    }

    return (
        <div className="login">
            <from>
                <div className='row'>
                    <div className='col'>
                        <button type="submit" className="btn btn-secondary" onClick={onSubmit_seleccion}>
                            Seleccion
                        </button>
                    </div>
                    <div className='col'>
                        <button type="submit" className="btn btn-primary" onClick={() => { navigate("/") }}>
                            Login
                        </button>
                    </div>
                </div>
            </from>
        </div>
    );
}