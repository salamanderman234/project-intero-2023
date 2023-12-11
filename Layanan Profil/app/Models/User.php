<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;


class User extends Model
{
    protected $hidden = ["created_at","updated_at","password","deleted_at","role","name"];
    public function profile() {
        $role = $this->role;
        $profile = [];
        switch($role) {
            case 'student': 
                return $this->hasOne(Student::class);
            break;
            case 'teacher': 
                return $this->hasOne(Teacher::class);
            break;
            default:
                return $this->hasOne(Admin::class);
            break;
        }
    }
    public function student()
    {
        return $this->hasOne(Student::class);
    }
    public function teacher()
    {
        return $this->hasOne(Teacher::class);
    }
}