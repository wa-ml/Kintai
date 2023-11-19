import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-reporter-register',
  templateUrl: './reporter-register.component.html',
  styleUrls: ['./reporter-register.component.scss']
})
export class ReporterRegisterComponent {
  constructor(private http: HttpClient) {}

  sendPostRequest(): void {
    this.http.post('/auth/kintaiLog', {}).subscribe(
      (response) => {
        console.log('POST request successful:', response);
        // ここに成功時の処理を記述
      },
      (error) => {
        console.error('Error in POST request:', error);
        // ここにエラーハンドリングを記述
      }
    );
  }
}
