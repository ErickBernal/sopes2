import { useNavigate } from 'react-router-dom'
import React, { useState, useEffect } from "react"
import axios from 'axios'



export default function Seleccion() {
    const navigate = useNavigate()


    const onSubmit_ram = () => {
        navigate('/ram')
    }
    const onSubmit_usb = () => {
        navigate('/usb')
    }


    return (
        <div className="login">
            <from>
                <div className='row'>
                    <div className='col'>
                        <button type="submit" className="btn btn-primary" onClick={onSubmit_ram}>
                            RAM
                        </button>
                    </div>
                    <div className='col'>
                        <button type="submit" className="btn btn-warning" onClick={onSubmit_usb}>
                            USB
                        </button>
                    </div>
                </div>
            </from>
        </div>
    );
}