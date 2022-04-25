import { NgModule ,CUSTOM_ELEMENTS_SCHEMA} from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { SidebarModule, MenuAllModule, TreeViewAllModule} from '@syncfusion/ej2-angular-navigations';
import { ListViewAllModule } from '@syncfusion/ej2-angular-lists';
import { RadioButtonModule, ButtonModule } from '@syncfusion/ej2-angular-buttons';
import { DropDownListModule } from '@syncfusion/ej2-angular-dropdowns';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { environment } from 'src/environments/environment';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import { HttpClientModule } from '@angular/common/http';
import {NgxMaterialTimepickerModule} from 'ngx-material-timepicker';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'; 
import { CoursesService } from './courses/courses.service';


// import the ScheduleModule for the Schedule component
import { ScheduleModule } from '@syncfusion/ej2-angular-schedule';
import { AppComponent } from './app.component';
import { SidebarMenuComponent } from './sidebar-menu/sidebar-menu.component';
import { ScheduleComponent } from './schedule/schedule.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { HomepageComponent } from './homepage/homepage.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import {RouterModule} from '@angular/router';
import { AvailableRoutes } from './app.routing';
import { UserProfileComponent } from './user-profile/user-profile.component';
import { AdminUsersComponent } from './admin/admin-users/admin-users.component';
import { AdminCoursesComponent } from './admin/admin-courses/admin-courses.component';
import { CoursesComponent } from './courses/courses.component';
import { NavbarComponent } from './navbar/navbar.component';
import { CreateComponent } from './create/create.component';

@NgModule({
  //declaration of ej2-angular-schedule module into NgModule
  declarations: [
    AppComponent,
    CoursesComponent,
    CreateComponent,
    SidebarMenuComponent,
    ScheduleComponent,
    LoginComponent,
    RegisterComponent,
    HomepageComponent,
    ToolbarComponent,
    UserProfileComponent,
    AdminUsersComponent,
    AdminCoursesComponent,
    NavbarComponent
  ],
  imports: [
    NgbModule,
    HttpClientModule,
    NgxMaterialTimepickerModule,
    BrowserAnimationsModule,
    SidebarModule, 
    MenuAllModule, 
    RadioButtonModule,
    DropDownListModule, 
    ButtonModule, 
    TreeViewAllModule, 
    ListViewAllModule,
    BrowserModule,
    ScheduleModule,
    FormsModule,
    ReactiveFormsModule,
    RouterModule,
    RouterModule.forRoot(AvailableRoutes)
  ],
  schemas: [
    CUSTOM_ELEMENTS_SCHEMA
  ],
  providers: [
    CoursesService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
