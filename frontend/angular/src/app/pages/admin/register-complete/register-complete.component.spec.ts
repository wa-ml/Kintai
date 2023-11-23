import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterCompleteComponent } from './register-complete.component';

describe('RegisterCompleteComponent', () => {
  let component: RegisterCompleteComponent;
  let fixture: ComponentFixture<RegisterCompleteComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [RegisterCompleteComponent]
    });
    fixture = TestBed.createComponent(RegisterCompleteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
