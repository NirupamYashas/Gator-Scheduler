import { ScheduleComponent } from '@syncfusion/ej2-angular-schedule';
import { SidebarMenuComponent } from './sidebar-menu/sidebar-menu.component';
import { HomepageComponent } from "./homepage/homepage.component";
import { RegisterComponent } from "./register/register.component";
import { LoginComponent } from "./login/login.component";
import { UserProfileComponent } from "./user-profile/user-profile.component";
import { AdminAuthGuard } from "./_helpers/admin-auth.guard";
import { AuthGuard } from './_helpers/auth.guard';
import { AdminUsersComponent } from "./admin/admin-users/admin-users.component";
import { AdminCoursesComponent } from './admin/admin-courses/admin-courses.component';
import { CoursesComponent } from "./courses/courses.component";
import { CreateComponent } from "./create/create.component";

export const AvailableRoutes: any = [
    { path: "", component: HomepageComponent},
    { path: "register", component: RegisterComponent},
    { path: "login", component: LoginComponent},
    {path: "schedule", component: ScheduleComponent},
    {path: "timetable", component: SidebarMenuComponent},
    { path: "create", component: CreateComponent, canActivate:[AuthGuard] },
    { path: "courses", component: CoursesComponent,canActivate:[AuthGuard]},
    { path: "user-profile", component: UserProfileComponent, canActivate:[AuthGuard]},
    { path: "admin/users", component: AdminUsersComponent, canActivate:[AuthGuard,AdminAuthGuard]},
    { path: "admin/courses", component: AdminCoursesComponent, canActivate:[AuthGuard,AdminAuthGuard]}
];