import { CoursesService } from './courses.service';
import { Router,ActivatedRoute } from "@angular/router";

import {  Component , OnInit} from '@angular/core';
import { Course } from '../_models/course';

@Component({
    selector: 'projects',
    templateUrl: './courses.component.html',
    styleUrls: ['./courses.component.css']
})
export class CoursesComponent {
    title = "List Of Courses";
    
    constructor(public service: CoursesService,private router: Router,
        private route: ActivatedRoute,){}
    
    courses: any;
    ngOnInit(): void{

        this.service.getCourses().subscribe(data => {
            this.courses = data;
        })  
    }

    editProject(course: Course) {
        let route = '/courses/course-page';
        this.router.navigate([route], { queryParams: { id: course.ID } });
    }
}