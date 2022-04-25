import { Component, OnInit,ViewEncapsulation, Inject, ViewChild } from '@angular/core';
import { SidebarComponent } from '@syncfusion/ej2-angular-navigations';
import { Menu, MenuItemModel } from '@syncfusion/ej2-navigations';
import { enableRipple } from '@syncfusion/ej2-base';
import { CoursesService } from '../courses/courses.service';
// import { checkboxdata } from './dataSource';


@Component({
  selector: 'app-sidebar-menu',
  templateUrl: './sidebar-menu.component.html',
  styleUrls: ['./sidebar-menu.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class SidebarMenuComponent{

courses: any;

  @ViewChild('sidebarMenuInstance')
  public sidebarMenuInstance!: SidebarComponent;
  public width: string = '220px';
  public mediaQuery: string = ('(min-width: 600px)');
  public target: string = '.main-content';
  public dockSize: string = '0px'
  public enableDock: boolean = true;
  constructor(public service: CoursesService) {
      
  }

  ngOnInit(): void{

    this.service.getCourses().subscribe(data => {
        this.courses = data;
    })  
  }

  //Define an array of JSON data
  public data: Object[] = ['course-1','course-2','course-3','course-4','course-5'];


  public menuItems: MenuItemModel[] = [
      {
          text: 'Overview',
          iconCss: 'icon-globe icon',
          items: [
              { text: 'All Data' },
              { text: 'Category2' },
              { text: 'Category3' }
          ]
      },
      {
          text: 'Notification',
          iconCss: 'icon-bell-alt icon',
          items: [
              { text: 'Message' },
              { text: 'Facebook' },
              { text: 'Twitter' }
          ]
      },
      {
          text: 'Comments',
          iconCss: 'icon-comment-inv-alt2 icon',
          items: [
              { text: 'Category1' },
              { text: 'Category2' },
              { text: 'Category3' }
          ]
      },
      {
          text: 'Bookmarks',
          iconCss: 'icon-bookmark icon',
          items: [
              { text: 'All Comments' },
              { text: 'Add Comments' },
              { text: 'Delete Comments' }
          ]
      },
      {
          text: 'Images',
          iconCss: 'icon-picture icon',
          items: [
              { text: 'Add Name' },
              { text: 'Add Mobile Number' },
              { text: 'Add Imaage' },
          ]
      },
      {
          text: 'Users ',
          iconCss: 'icon-user icon',
          items: [
              { text: 'Mobile1' },
              { text: 'Mobile2' },
              { text: 'Telephone' }
          ]
      },
      {
          text: 'Settings',
          iconCss: 'icon-eye icon',
          items: [
              { text: 'Change Profile' },
              { text: 'Add Name' },
              { text: 'Add Details' }
          ]
      },
      {
          text: 'Info',
          iconCss: 'icon-tag icon',
          items: [
              { text: 'Facebook' },
              { text: 'Mobile' },
          ]
      }
  ];
   public AccountMenuItem: MenuItemModel[] = [
      {
          text: 'Account',
          items: [
              { text: 'Profile' },
              { text: 'Sign out' },
          ]
      }
  ];
  // open new tab
  // newTabClick(): void {
  //     let URL = location.href.replace(location.search,'');
  //     document.getElementById('newTab').setAttribute('href', URL.split('#')[0] + 'sidebar/sidebar-menu');
  // }

  openClick() {
      this.sidebarMenuInstance.toggle();
  }
};
