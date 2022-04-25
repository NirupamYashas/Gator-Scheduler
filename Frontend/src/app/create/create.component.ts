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
      course_name: new FormControl('', [Validators.required]),
      course_code: new FormControl('', [Validators.required]),
      department_name: new FormControl('', [Validators.required]),
      course_instructor: new FormControl('', [Validators.required]),
      course_days: this.fb.array(this.checkboxes.map(x => false),[Validators.required]),
      // course_days: new FormArray(this.checkboxes.map(x => false), [Validators.required]),
      monday_starting_time: new FormControl('', [Validators.required]),
      monday_ending_time: new FormControl('', [Validators.required]),
      tuesday_starting_time: new FormControl('', [Validators.required]),
      tuesday_ending_time: new FormControl('', [Validators.required]),
      wednesday_starting_time: new FormControl('', [Validators.required]),
      wednesday_ending_time: new FormControl('', [Validators.required]),
      thursday_starting_time: new FormControl('', [Validators.required]),
      thursday_ending_time: new FormControl('', [Validators.required]),
      friday_starting_time: new FormControl('', [Validators.required]),
      friday_ending_time: new FormControl('', [Validators.required]),
    })

    const checkboxControl = (this.addCourseForm.controls['course_days'] as FormArray);
    this.subscription = checkboxControl.valueChanges.subscribe(checkbox => {
      checkboxControl.setValue(
        checkboxControl.value.map((value:any, i:any) => value ? this.checkboxes[i].value : false),
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
    // this.coursesService.createCourse(formData).subscribe(data => {})
  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }

}
