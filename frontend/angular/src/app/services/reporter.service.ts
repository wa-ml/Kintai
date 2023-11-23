import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { Observable, throwError } from 'rxjs';
import { Reporter } from '../types/reporter';

@Injectable({
  providedIn: 'root'
})
export class ReporterService {
  // TODO: envファイルなどに定義してエンドポイントを本番と開発で切り替えられるようにする
  private reportersUrl = 'http://localhost:8080/auth/users';
  private reporterUrl = 'http://localhost:8080/auth/user';
  private checkUrl = 'http://localhost:8080/auth/check_kintai_status';
  private kintaiLogUrl = 'http://localhost:8080/auth/kintaiLog';
  private headers = {
    headers: new HttpHeaders().set('Authorization', `Bearer ${localStorage.getItem('token')}`)
  }
  
  constructor(private http: HttpClient ) {}

  getReporters(): Observable<Reporter[]> {
    return this.http.get<Reporter[]>(this.reportersUrl, this.headers)
               .pipe(catchError(e => {
                console.log("error:", e);
                return throwError(() => e);
               }));
  }

  postReporter(name: string, email: string): Observable<any> {
    return this.http.post<any>(this.reporterUrl, {
      "Name": name,
      "Email": email,
      "Password": "hogehoge"
    }, this.headers)
    .pipe(catchError(e => {
      console.log("error:", e);
      return throwError(() => e);
    }))
  }

  checkStatus(): Observable<string> {
    return this.http.get<string>(this.checkUrl, this.headers)
    .pipe(catchError(e => {
     console.log("error:", e);
     return throwError(() => e);
    }));
  }

  postKintai(): Observable<any> {
    return this.http.post<any>(this.kintaiLogUrl, {}, this.headers)
    .pipe(catchError(e => {
      console.log("error:", e);
      return throwError(() => e);
    }))
  }
}
