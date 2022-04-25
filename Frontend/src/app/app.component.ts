import { Component } from '@angular/core';
import { EventSettingsModel, DayService, WeekService, WorkWeekService, MonthService, AgendaService ,View } from '@syncfusion/ej2-angular-schedule';
import { AuthenticationService } from './services/authentication.service';
import { Router } from '@angular/router';
import { UsersService } from './services/users.service';
import { LocationStrategy, PlatformLocation, Location } from '@angular/common';

@Component({
  selector: 'app-root',
  providers: [DayService, WeekService, WorkWeekService, MonthService, AgendaService],
  // template: `<ejs-schedule width='100%' height='100%' [selectedDate]='selectedDate'
  // [eventSettings]='eventSettings'> </ejs-schedule>`,
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Gator-Scheduler';

  constructor(public location: Location,private userService:  UsersService, private auth: AuthenticationService, router: Router){
    // auth.user$.subscribe(user => {
    //   if(user){

    //     let returnUrl: string | any = localStorage.getItem('returnUrl');
    //     router.navigateByUrl(returnUrl);
    //   }
    // })
  }

  removeNavbar() {
    var title = this.location.prepareExternalUrl(this.location.path());
    title = title.slice(1);
    if(title === 'courses'){
        return true;
    }
    else {
        return false;
    }
  }  

  public setView: View = 'Month';

  // public data: object [] = [{
  //       id: 2,
  //       eventName: 'Meeting',
  //       startTime: new Date(2018, 1, 15, 10, 0),
  //       endTime: new Date(2018, 1, 15, 12, 30),
  //       isAllDay: false
  // }];
  // public selectedDate: Date = new Date(2018, 1, 15);
  // public eventSettings: EventSettingsModel = {dataSource: this.data,
  //   fields: {
  //     id: 'id',
  //     subject: { name: 'eventName' },
  //     isAllDay: { name: 'isAllDay' },
  //     startTime: { name: 'startTime' },
  //     endTime: { name: 'endTime' },
  //   }
  // };
     
}
