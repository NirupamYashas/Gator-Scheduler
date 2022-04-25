import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormArray, FormControl, FormGroup ,Validators } from '@angular/forms';
import { CoursesService } from '../courses/courses.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {

  addCourseForm! : FormGroup;
  subscription: Subscription;
  
  
  checkboxes = [{
    name: 'Monday',
    value: 'monday'
  }, {
    name: 'Tuesday',
    value: 'tuesday'
  }, {
    name: 'Wednesday',
    value: 'wednesday'
  }, {
    name: 'Thursday',
    value: 'thursday'
  }, {
    name: 'Friday',
    value: 'friday'
  }];

  constructor(public coursesService: CoursesService,private fb: FormBuilder) { }

  msgTrue = false;
  ngOnInit(): void {

    this.addCourseForm = this.fb.group({
      name: new FormControl('', [Validators.required]),
      code: new FormControl('', [Validators.required]),
      department: new FormControl('', [Validators.required]),
      instructor: new FormControl('', [Validators.required]),
      days: this.fb.array(this.checkboxes.map(x => false),[Validators.required]),
      // course_days: new FormArray(this.checkboxes.map(x => false), [Validators.required]),
      monday_start_time: new FormControl('', [Validators.required]),
      monday_end_time: new FormControl('', [Validators.required]),
      tuesday_start_time: new FormControl('', [Validators.required]),
      tuesday_end_time: new FormControl('', [Validators.required]),
      wednesday_start_time: new FormControl('', [Validators.required]),
      wednesday_end_time: new FormControl('', [Validators.required]),
      thursday_start_time: new FormControl('', [Validators.required]),
      thursday_end_time: new FormControl('', [Validators.required]),
      friday_start_time: new FormControl('', [Validators.required]),
      friday_end_time: new FormControl('', [Validators.required]),
    })

    const checkboxControl = (this.addCourseForm.controls['days'] as FormArray);
    this.subscription = checkboxControl.valueChanges.subscribe(checkbox => {
      checkboxControl.setValue(
        checkboxControl.value.map((value:any, i:any) => value ? true : false),
        { emitEvent: false }
      );
    });
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

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }

}
