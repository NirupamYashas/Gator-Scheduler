import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup ,Validators } from '@angular/forms';
import { CoursesService } from '../courses/courses.service';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {

  addCourseForm! : FormGroup; 

  constructor(public coursesService: CoursesService) { }

  msgTrue = false;
  ngOnInit(): void {

    this.addCourseForm = new FormGroup({
      course_name: new FormControl('', [Validators.required]),
      course_code: new FormControl('', [Validators.required]),
      department_name: new FormControl('', [Validators.required]),
      course_instructor: new FormControl('', [Validators.required]),
      course_days: new FormControl('', [Validators.required]),
      starting_time: new FormControl('', [Validators.required]),
      ending_time: new FormControl('', [Validators.required]),
    })
  }

  addCourseSubmit(formData: any){
     console.log(formData);
    //  var projectName = this.addCourseForm.getRawValue().project_name;
    //  var departmentName = this.addProjectForm.getRawValue().department_name;
    //  var ufMail = this.addProjectForm.getRawValue().uf_mail;
    //  var githubLink = this.addProjectForm.getRawValue().github_link;
    //  const newFormData = { name: projectName, department: departmentName, email: ufMail , link: githubLink };
    this.coursesService.createCourse(formData).subscribe(data => {})
  }

}
