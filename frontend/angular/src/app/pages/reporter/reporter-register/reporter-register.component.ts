import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ReporterService } from 'src/app/services/reporter.service';

@Component({
  selector: 'app-reporter-register',
  templateUrl: './reporter-register.component.html',
  styleUrls: ['./reporter-register.component.scss']
})
export class ReporterRegisterComponent implements OnInit {
  isActive = true;
  kintaiLog = [];
  constructor(private http: HttpClient, private reporterService: ReporterService) {}

  ngOnInit(): void {
    this.checkStatus()
  }

  sendPostRequest(): void {
    this.reporterService.postKintai().subscribe((kintaiLog) => {
      this.isActive = kintaiLog.Status == "Active";
      this.kintaiLog = kintaiLog;
    })
  }
  checkStatus(): void {
    this.reporterService.checkStatus().subscribe((isActive) => {
      this.isActive = isActive === "Active";
    });
  }
}
