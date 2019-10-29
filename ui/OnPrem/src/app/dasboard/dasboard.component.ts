import { Component, OnInit } from '@angular/core';
import * as Highcharts from 'highcharts';
import {GraphService} from 'service/graph.service';

declare var require: any;
const Boost = require('highcharts/modules/boost');
const noData = require('highcharts/modules/no-data-to-display');
const More = require('highcharts/highcharts-more');
const Gauge = require('highcharts/modules/solid-gauge');

Boost(Highcharts);
noData(Highcharts);
More(Highcharts);
noData(Highcharts);
Gauge(Highcharts);

@Component({
  selector: 'Dashboard',
  templateUrl: './dasboard.component.html',
  styleUrls: ['./dasboard.component.css']
})
export class DasboardComponent implements OnInit {

  constructor(private gService: GraphService) { }

  currentStatus;
  currentTemp;

  initGraphs(){
    const option = this.gService.getCurrentStatusData();
    const tempGaugeOpt = this.gService.getCurrentTemeratureData();
    // @ts-ignore
    tempGaugeOpt.yAxis.max = 50;
    // @ts-ignore
    tempGaugeOpt.yAxis.min = 10;
    // @ts-ignore
    this.currentStatus = Highcharts.chart('currentstatus', option);
    // @ts-ignore
    this.currentTemp = Highcharts.chart('currentTemperature', tempGaugeOpt);
  }

  initListners(){
    let point = this.currentStatus.series[0].points[0];
    const tempPoint = this.currentTemp.series[0].points[0];
    this.gService.getDashboardData().subscribe(data => {
      const option = this.gService.getCurrentStatusData();
      // @ts-ignore
      if (option.yAxis.max !== data.data.CState.ON + data.data.CState.OFF){
        // @ts-ignore
        option.yAxis.max = data.data.CState.ON + data.data.CState.OFF
        // @ts-ignore
        this.currentStatus = Highcharts.chart('currentstatus', option);
        point = this.currentStatus.series[0].points[0];
      }
      // @ts-ignore
      point.update(data.data.CState.ON);

      // @ts-ignore
      tempPoint.update(data.data.Temperature);
    });
  }

  ngOnInit() {
    this.initGraphs();
    this.initListners();
  }

}
