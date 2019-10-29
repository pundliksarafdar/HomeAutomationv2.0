import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class HistoryService {
  private value: number;

  constructor(private http: HttpClient) { }

  getHistoryData() {
    return this.http.get('/rest/analytics/history');
  }

  getHistoryChartOption(seriesOptions, label, valMin, valMax, valLabel) {
    return {
      title: {
        text: label
      },
      rangeSelector: {
        selected: 4
      },
      xAxis: {
        type: 'datetime'
      },
      yAxis: {
        allowDecimals: false,
        min: valMin,
        max: valMax,
        title: {
          text: valLabel
        },
        plotLines: [{
          value: 0,
          width: 2,
          color: 'silver'
        }],
        labels: {
          formatter : (e) => {
            if (e.value === 0){
              return 'OFF';
            } else if(e.value === 1){
              return 'ON';
            } else {
              return  e.value;
            }
          }
        }
      },

      plotOptions: {
        series: {
          compare: 'percent',
            showInNavigator: true
        }
      },

      tooltip: {
        pointFormat: '<span style="color:{series.color}">{series.name}</span>: <b>{point.y}</b><br/>',
          valueDecimals: 0,
          split: true
      },

      series: seriesOptions
    };
  }
}
