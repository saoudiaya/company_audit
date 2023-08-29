import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SigninComponent } from './signin/signin.component';
import { SignupComponent } from './signup/signup.component';
import { AccueilComponent } from './accueil/accueil.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { TeamsComponent } from './teams/teams.component';
import { CompaniesComponent } from './companies/companies.component';
import { AuditComponent } from './audit/audit.component';
import { ControllistComponent } from './controllist/controllist.component';
import { DocumentsComponent } from './documents/documents.component';
import { MeetingComponent } from './meeting/meeting.component';
import { TasksComponent } from './tasks/tasks.component';
import { PermissionComponent } from './permission/permission.component';
import { RequirementComponent } from './requirement/requirement.component';
import { PerimeterComponent } from './perimeter/perimeter.component';
import { ObjectiveComponent } from './objective/objective.component';
import { CriteriaComponent } from './criteria/criteria.component';
import { NormComponent } from './norm/norm.component';



const routes: Routes = [
  { path: '', redirectTo: '/signin', pathMatch: 'full' }, // Redirect to /signin
  { path: 'signin', component: SigninComponent },
  { path: 'signup', component: SignupComponent },
  { path: 'accueil', component: AccueilComponent},
  { path: 'dashboard', component: DashboardComponent},
  { path: 'teams', component: TeamsComponent},
  { path: 'companies', component: CompaniesComponent},
  { path: 'audit', component: AuditComponent},
  { path: 'controllist', component: ControllistComponent},
  { path: 'documents', component: DocumentsComponent},
  { path: 'meeting', component: MeetingComponent},
  { path: 'tasks', component: TasksComponent},
  { path: 'permission', component: PermissionComponent},
  { path: 'requirement', component: RequirementComponent},
  { path: 'perimeter', component: PerimeterComponent},
  { path: 'objective', component: ObjectiveComponent},
  { path: 'criteria', component: CriteriaComponent},
  { path: 'norm', component: NormComponent},
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
