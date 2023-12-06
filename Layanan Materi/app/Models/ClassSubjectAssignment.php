<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class ClassSubjectAssignment extends Model
{
    use HasFactory, SoftDeletes;
    protected $guarded = ["id"];

    public function class_subject()
    {
        return $this->belongsTo(ClassSubject::class,'class_subject_id', 'id');
    }
}
