import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import { RobotApi, TablesLLModel, SettingUpdateParam } from '../../openapi/wrobot';


@Component({
    selector: 'page-setting-update',
    templateUrl: 'update.html',
    styleUrls: ['update.scss']
})
export class SettingUpdateComponent implements OnInit {

    public formdata = {} as SettingUpdateParam;
    public llmodels: Array<TablesLLModel> = [];

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) {
        this.getLLModels();
    }

    public ngOnInit() {
        const name = this.route.snapshot.paramMap.get('name');
        name && this.getSetting(name);
    }

    public getSetting(name: string) {
        RobotApi.settingDetail({ name }).then((data) => {
            data && Object.assign(this.formdata, data);
        });
    }

    public updateSetting() {
        this.formdata.value = String(this.formdata.value);
        RobotApi.settingUpdate(this.formdata).then(() => {
            this.router.navigate(['setting/list']);
        });
    }

    public getLLModels() {
        RobotApi.llmodelList({}).then((data) => {
            this.llmodels = data || [];
        });
    }
}