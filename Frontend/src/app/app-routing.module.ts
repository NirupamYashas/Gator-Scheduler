import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ScheduleComponent } from '@syncfusion/ej2-angular-schedule';
import { SidebarMenuComponent } from './sidebar-menu/sidebar-menu.component';

const routes: Routes = [
  {path: "schedule", component: ScheduleComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
