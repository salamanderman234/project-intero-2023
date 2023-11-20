<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;


class User extends Model
{
    public function student()
    {
        return $this->hasOne(Student::class);
    }
    public function teacher()
    {
        return $this->hasOne(Teacher::class);
    }
}