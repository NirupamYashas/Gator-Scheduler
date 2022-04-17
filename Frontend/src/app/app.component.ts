import { Component } from '@angular/core';
import { EventSettingsModel, DayService, WeekService, WorkWeekService, MonthService, AgendaService ,View } from '@syncfusion/ej2-angular-schedule';

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
