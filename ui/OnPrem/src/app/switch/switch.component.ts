import {Component, OnInit} from '@angular/core';
import {SwitchServiceService} from '../service/switch-service.service';

@Component({
  selector: 'Switch',
  templateUrl: './switch.component.html',
  styleUrls: ['./switch.component.css']
})
export class SwitchComponent implements OnInit {

  switchData = [];
  selectedVal;

  constructor(private switchSeervice: SwitchServiceService) {
  }

  public loadData() {
    this.switchSeervice.getSwitchDetails().subscribe(data => {
      // @ts-ignore
      const dataSwitch = data.data;
      // @ts-ignore
      const dataPort = data.port;
      // tslint:disable-next-line:forin
      for (let ip in dataSwitch) {
        const dataForUi = [];

        let switchName = '';
        for (let key in dataSwitch[ip]) {
          if (key !== 'switchName') {
            dataForUi.push({name: key, status: dataSwitch[ip][key], port: dataPort[ip][key]});
          } else {
            switchName = dataSwitch[ip][key];
          }

        }
        this.switchData.push({dataForUi, ip, switchName});
      }
    });
  }

  public onSwitchChange(ip, port, currentState) {
    console.log(ip, port, currentState);
    const status = currentState.toString().toLowerCase();
    const portData = {ip, port, status };
    this.switchSeervice.postSwitchDetails(portData);
  }

  ngOnInit() {
    this.loadData();
  }

}
