import { NgModule ,CUSTOM_ELEMENTS_SCHEMA} from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { SidebarModule, MenuAllModule, TreeViewAllModule} from '@syncfusion/ej2-angular-navigations';
import { ListViewAllModule } from '@syncfusion/ej2-angular-lists';
import { RadioButtonModule, ButtonModule } from '@syncfusion/ej2-angular-buttons';
import { DropDownListModule } from '@syncfusion/ej2-angular-dropdowns';


// import the ScheduleModule for the Schedule component
import { ScheduleModule } from '@syncfusion/ej2-angular-schedule';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SidebarMenuComponent } from './sidebar-menu/sidebar-menu.component';
import { ScheduleComponent } from './schedule/schedule.component';

@NgModule({
  //declaration of ej2-angular-schedule module into NgModule
  declarations: [
    AppComponent,
    SidebarMenuComponent,
    ScheduleComponent
  ],
  imports: [
    SidebarModule, 
    MenuAllModule, 
    RadioButtonModule,
    DropDownListModule, 
    ButtonModule, 
    TreeViewAllModule, 
    ListViewAllModule,
    BrowserModule,
    AppRoutingModule,
    ScheduleModule
  ],
  schemas: [
    CUSTOM_ELEMENTS_SCHEMA
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
