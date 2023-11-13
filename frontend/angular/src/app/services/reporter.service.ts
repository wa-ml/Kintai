import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { Observable, throwError } from 'rxjs';
import { Reporter } from '../types/reporter';

@Injectable({
  providedIn: 'root'
})
export class ReporterService {
  // TODO: envファイルなどに定義してエンドポイントを本番と開発で切り替えられるようにする
  private attendancesUrl = 'http://localhost:3000/reporters';

  constructor(private http: HttpClient) {}

  getReporters(): Observable<Reporter[]> {
    return this.http.get<Reporter[]>(this.attendancesUrl, {})
               .pipe(catchError(e => {
                console.log("error:", e);
                return throwError(() => e);
               }));
  }
}
