import { Component, OnInit } from '@angular/core';
import { Chart, registerables } from 'chart.js';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  ngOnInit(): void {
    Chart.register(...registerables);
    this.createBarChart();
    this.createPieChart();
  }

  createBarChart() {
    const ctx = document.getElementById('auditChart') as HTMLCanvasElement;
    const chart = new Chart(ctx, {
      type: 'bar',
      data: {
        labels: ['2019', '2020', '2021', '2022', '2023'],
        datasets: [
          {
            label: 'Number of Audits',
            data: [50, 100, 80, 120, 90],
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
            borderColor: 'rgba(75, 192, 192, 1)',
            borderWidth: 1
          }
        ]
      },
      options: {
        scales: {
          y: {
            beginAtZero: true,
          }
        }
      }
    });
  }
  createPieChart() {
    const ctx = document.getElementById('auditChart1') as HTMLCanvasElement;
    const chart = new Chart(ctx, {
      type: 'pie',
      data: {
        labels: ['Category A', 'Category B', 'Category C', 'Category D'],
        datasets: [
          {
            label: 'Audits per Category',
            data: [20, 30, 15, 10],
            backgroundColor: ['rgba(255, 99, 132, 0.8)', 'rgba(54, 162, 235, 0.8)', 'rgba(255, 205, 86, 0.8)', 'rgba(75, 192, 192, 0.8)'],
            borderColor: 'rgba(255, 255, 255, 1)',
            borderWidth: 2
          }
        ]
      }
    });
  }
}
