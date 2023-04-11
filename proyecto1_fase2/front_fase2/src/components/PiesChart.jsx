import React from 'react';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Pie } from 'react-chartjs-2';

ChartJS.register(ArcElement, Tooltip, Legend);

export default function Pies({ total, disponible, en_uso }) {
    var options = {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
            tooltip: {
                callbacks: {
                    label: function (context) {
                        var label = context.label || '';
                        var value = context.parsed || 0;
                        var percent = context.dataset.data.reduce(function (a, b) { return a + b; }, 0) > 0 ? value * 100 / context.dataset.data.reduce(function (a, b) { return a + b; }, 0) : 0;
                        percent = Math.round(percent * 100) / 100;
                        return  percent + '%';
                    }
                }
            }
        }
    };

    var data = {
        labels: ['En_uso', 'Disponible'],
        datasets: [
            {
                label: 'Cantidad de Memoria Ram',
                data: [en_uso, disponible],
                backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(255, 206, 86, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(255, 206, 86, 1)',
                    'rgba(54, 162, 235, 1)',
                    'rgba(75, 192, 192, 1)',
                    'rgba(153, 102, 255, 1)',
                ],
                borderWidth: 2,
            },
        ],
    };

    return <Pie data={data} options={options} />;
}
