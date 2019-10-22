import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class SwitchServiceService {

  constructor(private http: HttpClient) {
  }

  public getSwitchDetails() {
    return this.http.get('/rest/switch/info');
  }

  public postSwitchDetails(portdata) {
    this.http.post('/rest/switch/toggle', portdata).subscribe(
      data  => {
        console.log('POST Request is successful ', data);
      },
      error  => {
        console.log('Error', error);
      }

    );
  }
}
