import { Component, OnInit } from '@angular/core';
import { EventSettingsModel, DayService, WeekService, WorkWeekService, MonthService, AgendaService, View } from '@syncfusion/ej2-angular-schedule';

@Component({
  selector: 'app-schedule',
  providers: [DayService, WeekService, WorkWeekService, MonthService, AgendaService],
  templateUrl: './schedule.component.html',
  styleUrls: ['./schedule.component.css']
})
export class ScheduleComponent{
  title = 'Gator-Scheduler';

  // public setView: View = 'Month';

  public data: object [] = [{
    id: 2,
    eventName: 'Meeting',
    startTime: new Date(2022, 3, 22, 10, 0),
    endTime: new Date(2022, 3, 22, 12, 30),
    isAllDay: false
  }];
  public selectedDate: Date = new Date(2022, 3, 22);
  public eventSettings: EventSettingsModel = {dataSource: this.data,
    fields: {
      id: 'id',
      subject: { name: 'eventName' },
      isAllDay: { name: 'isAllDay' },
      startTime: { name: 'startTime' },
      endTime: { name: 'endTime' },
    }
  };


}
