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
  private attendancesUrl = 'http://localhost:8080/auth/users';

  constructor(private http: HttpClient ) {}

  getReporters(): Observable<Reporter[]> {
    return this.http.get<Reporter[]>(this.attendancesUrl, {
      headers: new HttpHeaders().set('Authorization', `Bearer ${localStorage.getItem('token')}`)
    })
               .pipe(catchError(e => {
                console.log("error:", e);
                return throwError(() => e);
               }));
  }
}
