import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReporterRecordComponent } from './reporter-record.component';

describe('ReporterRecordComponent', () => {
  let component: ReporterRecordComponent;
  let fixture: ComponentFixture<ReporterRecordComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ReporterRecordComponent]
    });
    fixture = TestBed.createComponent(ReporterRecordComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
