import { TestBed } from '@angular/core/testing';

import { SwitchServiceService } from './switch-service.service';

describe('SwitchServiceService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: SwitchServiceService = TestBed.get(SwitchServiceService);
    expect(service).toBeTruthy();
  });
});
