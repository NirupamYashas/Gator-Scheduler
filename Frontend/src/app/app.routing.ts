import { ScheduleComponent } from '@syncfusion/ej2-angular-schedule';
import { SidebarMenuComponent } from './sidebar-menu/sidebar-menu.component';
import { HomepageComponent } from "./homepage/homepage.component";
import { RegisterComponent } from "./register/register.component";
import { LoginComponent } from "./login/login.component";

export const AvailableRoutes: any = [
    { path: "", component: HomepageComponent},
    { path: "register", component: RegisterComponent},
    { path: "login", component: LoginComponent},
    {path: "schedule", component: ScheduleComponent},
    {path: "timetable", component: SidebarMenuComponent}
];