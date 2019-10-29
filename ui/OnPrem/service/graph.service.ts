import { Injectable } from '@angular/core';
import * as Highcharts from 'highcharts';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class GraphService {

  constructor( private http: HttpClient) {
  }

  gaugeOptions = {

    chart: {
      type: 'solidgauge'
    },

    title: null,

    pane: {
      center: ['50%', '85%'],
      size: '140%',
      startAngle: -90,
      endAngle: 90,
      background: {
        backgroundColor:
          Highcharts.defaultOptions.legend.backgroundColor || '#EEE',
        innerRadius: '60%',
        outerRadius: '100%',
        shape: 'arc'
      }
    },

    tooltip: {
      enabled: false
    },

    // the value axis
    yAxis: {
      stops: [
        [0.1, '#55BF3B'], // green
        [0.5, '#DDDF0D'], // yellow
        [0.9, '#DF5353'] // red
      ],
      lineWidth: 0,
      minorTickInterval: null,
      tickAmount: 2,
      title: {
        y: -70
      },
      labels: {
        y: 16
      }
    },

    plotOptions: {
      solidgauge: {
        dataLabels: {
          y: 5,
          borderWidth: 0,
          useHTML: true
        }
      }
    }
  };

  getCurrentStatusData() {
    // @ts-ignore
    const chartSpeed = Highcharts.merge(this.gaugeOptions, {
      yAxis: {
        min: 0,
        max: 200,
        title: {
          text: 'Switches'
        }
      },

      credits: {
        enabled: false
      },

      series: [{
        name: 'Switches',
        data: [0],
        dataLabels: {
          format:
            '<div style="text-align:center">' +
            '<span style="font-size:25px">{y}</span><br/>' +
            '<span style="font-size:12px;opacity:0.4">On</span>' +
            '</div>'
        },
        tooltip: {
          valueSuffix: ' on'
        }
      }]

    });
    return chartSpeed;
  }

  getCurrentTemeratureData() {
    // @ts-ignore
    const chartSpeed = Highcharts.merge(this.gaugeOptions, {
      yAxis: {
        min: 0,
        max: 200,
        title: {
          text: 'Temperature'
        }
      },

      credits: {
        enabled: false
      },

      series: [{
        name: 'Temperature',
        data: [0],
        dataLabels: {
          format:
            '<div style="text-align:center">' +
            '<span style="font-size:25px">{y}</span><br/>' +
            '<span style="font-size:12px;opacity:0.4">C</span>' +
            '</div>'
        },
        tooltip: {
          valueSuffix: ' C'
        }
      }]

    });
    return chartSpeed;
  }

  getDashboardData() {
    return this.http.get('/rest/analytics/dashboard');
  }
}
