import { Component, OnInit } from '@angular/core';
import {HistoryService} from '../service/history.service';
import * as Highcharts from 'highcharts';

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
  selector: 'History',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.css']
})
export class HistoryComponent implements OnInit {

  constructor(public hData: HistoryService) { }

  showHistoryData() {
    this.hData.getHistoryData().subscribe((data) => {
      // @ts-ignore
      const option = this.hData.getHistoryChartOption(data.historicData);
      // @ts-ignore
      const toption = this.hData.getHistoryChartOption(data.tempdata);
      // @ts-ignore
      Highcharts.chart('historycontainer', option);
      // @ts-ignore
      Highcharts.chart('tempcontainer', toption);
    });
  }
  ngOnInit() {
    this.showHistoryData();
  }

}
