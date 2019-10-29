import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class HistoryService {

  constructor(private http: HttpClient) { }

  getHistoryData(){
    return this.http.get('/rest/analytics/history');
  }

  getHistoryChartOption(seriesOptions){
    return {

      rangeSelector: {
        selected: 4
      },
      xAxis: {
        type: 'datetime'
      },
      yAxis: {
        plotLines: [{
          value: 0,
          width: 2,
          color: 'silver'
        }]
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
